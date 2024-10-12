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
