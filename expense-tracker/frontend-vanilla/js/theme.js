// Theme management module
const theme = {
    // Get current theme
    get() {
        return localStorage.getItem('theme') || 'light';
    },

    // Set theme
    set(mode) {
        localStorage.setItem('theme', mode);
        this.apply();
    },

    // Toggle between light and dark
    toggle() {
        const current = this.get();
        this.set(current === 'light' ? 'dark' : 'light');
    },

    // Apply current theme to document
    apply() {
        const isDark = this.get() === 'dark';
        document.documentElement.classList.toggle('dark', isDark);
        
        // Update theme toggle button if it exists
        const themeToggle = document.getElementById('themeToggle');
        if (themeToggle) {
            const icon = themeToggle.querySelector('i');
            if (icon) {
                icon.className = isDark ? 'fas fa-sun' : 'fas fa-moon';
            }
        }

        // Update theme color meta tag
        const metaThemeColor = document.querySelector('meta[name="theme-color"]');
        if (metaThemeColor) {
            metaThemeColor.content = isDark ? '#1a1a1a' : '#ffffff';
        }
