import { expenses as expensesApi } from './api.js';
import { notify, formatCurrency, formatDate, requireAuth, chartUtils } from './utils.js';

// Initialize dashboard when document is ready
document.addEventListener('DOMContentLoaded', () => {
    if (!requireAuth()) return;
    initializeDashboard();
});

async function initializeDashboard() {
    await Promise.all([
        loadSummaryData(),
        loadExpenseCharts(),
        loadRecentExpenses()
    ]);
}

async function loadSummaryData() {
    try {
