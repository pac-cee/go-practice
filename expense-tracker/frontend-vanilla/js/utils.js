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
