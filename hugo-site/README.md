# README.md

# Hugo Project Structure for "formsaas-theme"

This README provides an overview of the directory structure for the Hugo project named "formsaas-theme". Each directory and file serves a specific purpose in organizing the content, assets, and layouts of the site.

## Directory Structure

- **hugo-site/**: The root directory of the Hugo project.

- **archetypes/**: Contains archetype files for content creation.
  - **default.md**: The default archetype for new content.

- **assets/**: Holds static assets for the site.
  - **css/**: Contains stylesheets.
    - **main.css**: The main stylesheet for the site.
  - **js/**: Contains JavaScript files.
    - **main.js**: The main JavaScript file for the site.

- **content/**: Contains the content for the site.
  - **posts/**: Subdirectory for blog posts.
    - **_index.md**: Index file for the posts section.
  - **_index.md**: Main index file for the content.

- **data/**: Contains data files.
  - **config.yaml**: Configuration data for the site.

- **layouts/**: Contains layout templates for rendering content.
  - **_default/**: Default layout templates.
    - **baseof.html**: Base layout template.
    - **list.html**: Layout for listing content.
    - **single.html**: Layout for single content pages.
  - **partials/**: Reusable partial templates.
    - **head.html**: Head section of the HTML.
    - **header.html**: Header section of the HTML.
    - **footer.html**: Footer section of the HTML.
  - **index.html**: Layout for the homepage.

- **static/**: Contains static files served directly.
  - **favicon.ico**: Favicon for the site.

- **themes/**: Contains themes for the site.
  - **formsaas-theme/**: Directory for the "formsaas-theme".
    - **archetypes/**: Archetype files specific to the theme.
      - **default.md**: Default archetype for new content in the theme.
    - **assets/**: Theme-specific static assets.
      - **css/**: Theme stylesheets.
        - **theme.css**: Main stylesheet for the theme.
      - **js/**: Theme JavaScript files.
        - **theme.js**: Main JavaScript file for the theme.
    - **layouts/**: Layout templates specific to the theme.
      - **_default/**: Default layout templates for the theme.
        - **baseof.html**: Base layout template for the theme.
        - **list.html**: Layout for listing content in the theme.
        - **single.html**: Layout for single content pages in the theme.
      - **partials/**: Reusable partial templates for the theme.
        - **head.html**: Head section of the theme's HTML.
        - **header.html**: Header section of the theme's HTML.
        - **footer.html**: Footer section of the theme's HTML.
      - **index.html**: Layout for the theme's homepage.
    - **static/**: Static files specific to the theme.
      - **theme-assets/**: Directory for theme-specific assets.
    - **theme.toml**: Configuration file for the theme.

- **config.toml**: Main configuration file for the Hugo site.

This structure is designed to keep the project organized and maintainable, allowing for easy updates and modifications as needed.