// Utility functions for the application

// Format currency based on user's locale and currency preference
export function formatCurrency(amount, currency = 'USD') {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: currency
    }).format(amount);
}

// Format date to local string
export function formatDate(date) {
    return new Date(date).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
}

// Show notification using Notyf
export const notify = {
    success(message) {
        const notyf = new Notyf({
            duration: 3000,
            position: { x: 'right', y: 'top' }
        });
        notyf.success(message);
    },
    error(message) {
        const notyf = new Notyf({
            duration: 3000,
            position: { x: 'right', y: 'top' }
        });
        notyf.error(message);
    }
};

// Check if user is authenticated
export function isAuthenticated() {
    const token = localStorage.getItem('token');
    return !!token;
}

// Redirect if not authenticated
export function requireAuth() {
    if (!isAuthenticated()) {
        window.location.href = '/pages/login.html';
        return false;
    }
    return true;
}

// Handle API errors
export function handleError(error) {
    console.error('Error:', error);
    notify.error(error.message || 'An unexpected error occurred');
}

// Theme management
export const theme = {
    toggle() {
        const isDark = localStorage.getItem('darkMode') === 'true';
        localStorage.setItem('darkMode', !isDark);
        this.apply();
    },

    apply() {
        const isDark = localStorage.getItem('darkMode') === 'true';
        document.body.classList.toggle('dark-theme', isDark);
        const themeIcon = document.querySelector('#theme-toggle i');
        if (themeIcon) {
            themeIcon.className = isDark ? 'fas fa-sun' : 'fas fa-moon';
        }
    },

    init() {
