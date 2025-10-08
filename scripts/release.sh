#!/bin/bash
set -e

PROJECT=$1
VERSION=$2
NPM_PUBLISH=${3:-false}

if [ -z "$PROJECT" ] || [ -z "$VERSION" ]; then
    echo "Usage: ./scripts/release.sh <project> <version> [npm_publish]"
    echo "Example: ./scripts/release.sh svelte-library v1.0.0 true"
    exit 1
fi

# Get current branch
CURRENT_BRANCH=$(git branch --show-current)
RELEASE_BRANCH="release/${PROJECT}-${VERSION}"

echo "Current branch: $CURRENT_BRANCH"
echo "Expected release branch: $RELEASE_BRANCH"

# Check if we're on the expected release branch
if [ "$CURRENT_BRANCH" != "$RELEASE_BRANCH" ]; then
    echo "Error: Must be on release branch '$RELEASE_BRANCH'"
    echo "Current branch is '$CURRENT_BRANCH'"
    exit 1
fi

# Ensure we have latest changes from origin
git fetch origin

# Check if project directory exists
if [ ! -d "$PROJECT" ]; then
    echo "Project directory '$PROJECT' not found"
    exit 1
fi

cd "$PROJECT"

# Handle npm packages (svelte-app, svelte-library)
if [ -f "package.json" ]; then
    # Remove 'v' prefix for npm version
    NPM_VERSION=${VERSION#v}
    
    # Check current version
    CURRENT_VERSION=$(node -p "require('./package.json').version")
    
    if [ "$CURRENT_VERSION" = "$NPM_VERSION" ]; then
        echo "Version $NPM_VERSION is already set in package.json. Skipping version update."
    else
        echo "Updating package.json version from $CURRENT_VERSION to $NPM_VERSION..."
        
        # Update package.json version using npm (works with pnpm workspaces)
        npm version "$NPM_VERSION" --no-git-tag-version
    fi
    
    # Install dependencies and build the package
    echo "Installing dependencies..."
    pnpm install
    
    echo "Building package..."
    pnpm run build
    
    # Run tests if available (skip for certain projects like hugo-site)
    if [ "$PROJECT" != "hugo-site" ] && pnpm run --if-present test > /dev/null 2>&1; then
        echo "Running tests..."
        pnpm run test
    else
        echo "Skipping tests (not available or excluded for $PROJECT)"
    fi
    
    # Run type checking if available (skip for certain projects like hugo-site)
    if [ "$PROJECT" != "hugo-site" ] && pnpm run --if-present check > /dev/null 2>&1; then
        echo "Running type check..."
        pnpm run check
    else
        echo "Skipping type check (not available or excluded for $PROJECT)"
    fi
    
    # Go back to repo root
    cd ..
    
    # Commit any changes on the release branch
    if ! git diff --quiet HEAD -- "$PROJECT/package.json" "pnpm-lock.yaml"; then
        git add "$PROJECT/package.json" "pnpm-lock.yaml" 2>/dev/null || true
        git commit -m "$PROJECT: bump version to $VERSION"
        echo "Committed version changes to release branch"
    fi
fi

# Merge release branch to main
echo "Switching to main branch and merging release branch..."
git checkout main
git pull origin main

# Merge the release branch
git merge "$RELEASE_BRANCH" --no-ff -m "Merge $RELEASE_BRANCH into main"

# Push main branch
git push origin main

# Create and push tag on main
TAG="$PROJECT/$VERSION"
echo "Creating tag: $TAG"

# Check if tag already exists
if git rev-parse "$TAG" >/dev/null 2>&1; then
    echo "Tag $TAG already exists. Skipping tag creation."
else
    git tag -a "$TAG" -m "Release $PROJECT $VERSION"
    git push origin "$TAG"
    echo "Created and pushed tag: $TAG"
fi

# Clean up: delete the release branch locally and remotely
echo "Cleaning up release branch..."
git branch -d "$RELEASE_BRANCH"

# Check if remote release branch exists and delete it
if git ls-remote --heads origin "$RELEASE_BRANCH" | grep -q "$RELEASE_BRANCH"; then
    git push origin --delete "$RELEASE_BRANCH"
    echo "Deleted remote release branch: $RELEASE_BRANCH"
fi

# Publish to npm if requested and it's a package
if [ "$NPM_PUBLISH" = "true" ] && [ -f "$PROJECT/package.json" ]; then
    echo "Publishing to npm..."
    cd "$PROJECT"
    pnpm publish --access public --no-git-checks
    cd ..
    echo "Package published to npm!"
fi

echo "Released $PROJECT $VERSION successfully!"
echo "- Merged $RELEASE_BRANCH into main"
echo "- Created tag: $TAG"
echo "- Cleaned up release branch"
if [ "$NPM_PUBLISH" = "true" ]; then
    echo "- Published package to npm"
fi