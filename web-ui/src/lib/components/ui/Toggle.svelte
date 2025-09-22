<script lang="ts">
    interface Props {
        checked?: boolean;
        disabled?: boolean;
        size?: 'sm' | 'md' | 'lg';
        base?: string;
        background?: string;
        border?: string;
        rounded?: string;
        classes?: string;
        name?: string;
        id?: string;
        'aria-label'?: string;
        'aria-describedby'?: string;
        onchange?: (checked: boolean) => void;
    }

    let {
        checked = $bindable(false),
        disabled = false,
        size = 'md',
        base = '',
        background = '',
        border = '',
        rounded = '',
        classes = '',
        name = '',
        id = '',
        'aria-label': ariaLabel = '',
        'aria-describedby': ariaDescribedby = '',
        onchange
    }: Props = $props();

    // Size configurations
    const sizeClasses = {
        sm: {
            container: 'w-9 h-5',
            thumb: 'w-4 h-4',
            translate: 'translate-x-4'
        },
        md: {
            container: 'w-11 h-6',
            thumb: 'w-5 h-5',
            translate: 'translate-x-5'
        },
        lg: {
            container: 'w-14 h-7',
            thumb: 'w-6 h-6',
            translate: 'translate-x-7'
        }
    };

    let containerClass = $derived(() => {
        const baseClasses = base || 'relative inline-flex items-center cursor-pointer transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 focus:ring-offset-white dark:focus:ring-offset-gray-900';
        
        const sizeClass = sizeClasses[size].container;
        
        // Background color based on state
        const bgClass = background || (checked 
            ? 'bg-primary-600 hover:bg-primary-700' 
            : 'bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600'
        );
        
        const borderClass = border || '';
        const roundedClass = rounded || 'rounded-full';
        const disabledClass = disabled ? 'opacity-50 cursor-not-allowed' : '';

        return [
            baseClasses,
            sizeClass,
            bgClass,
            borderClass,
            roundedClass,
            disabledClass,
            classes
        ]
            .filter(Boolean)
            .join(' ');
    });

    let thumbClass = $derived(() => {
        const baseThumbClasses = 'absolute left-0 inline-block bg-white rounded-full shadow transform transition-transform duration-200 ease-in-out';
        const sizeClass = sizeClasses[size].thumb;
        const translateClass = checked ? sizeClasses[size].translate : 'translate-x-0';

        return [baseThumbClasses, sizeClass, translateClass]
            .filter(Boolean)
            .join(' ');
    });

    function handleChange(event: Event) {
        if (disabled) return;
        
        const target = event.target as HTMLInputElement;
        checked = target.checked;
        onchange?.(checked);
    }

    function handleKeydown(event: KeyboardEvent) {
        if (disabled) return;
        
        if (event.key === ' ' || event.key === 'Enter') {
            event.preventDefault();
            checked = !checked;
            onchange?.(checked);
        }
    }
</script>

<label class="inline-flex items-center {disabled ? 'cursor-not-allowed' : 'cursor-pointer'}">
    <!-- Hidden checkbox for accessibility and form submission -->
    <input
        type="checkbox"
        class="sr-only"
        bind:checked
        {disabled}
        {name}
        {id}
        aria-label={ariaLabel}
        aria-describedby={ariaDescribedby}
        onchange={handleChange}
    />
    
    <!-- Toggle visual element -->
    <span
        class={containerClass()}
        tabindex={disabled ? -1 : 0}
        role="switch"
        aria-checked={checked}
        onkeydown={handleKeydown}
    >
        <!-- Thumb -->
        <span class={thumbClass()}></span>
    </span>
</label>