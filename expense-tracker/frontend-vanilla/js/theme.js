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

