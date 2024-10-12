import { settings as settingsApi } from './api.js';
import { notify, requireAuth } from './utils.js';

// Initialize when document is ready
document.addEventListener('DOMContentLoaded', () => {
    if (!requireAuth()) return;
    initializeSettings();
});

async function initializeSettings() {
    initializeThemeSettings();
    initializeNotificationSettings();
    initializeSpendingLimits();
    initializeDataManagement();
    await loadSettings();
}

function initializeThemeSettings() {
    const themeToggle = document.getElementById('themeToggle');
    const currentTheme = localStorage.getItem('theme') || 'light';
    
    if (themeToggle) {
        themeToggle.checked = currentTheme === 'dark';
        themeToggle.addEventListener('change', (e) => {
            const newTheme = e.target.checked ? 'dark' : 'light';
            localStorage.setItem('theme', newTheme);
            document.documentElement.classList.toggle('dark', e.target.checked);
        });
    }
}

function initializeNotificationSettings() {
    const emailToggle = document.getElementById('emailNotifications');
    const pushToggle = document.getElementById('pushNotifications');
    const weeklyToggle = document.getElementById('weeklyReports');
    const monthlyToggle = document.getElementById('monthlyReports');

    const toggles = [emailToggle, pushToggle, weeklyToggle, monthlyToggle];
