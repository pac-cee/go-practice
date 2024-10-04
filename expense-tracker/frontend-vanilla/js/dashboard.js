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
        const response = await expensesApi.getSummary();
        updateSummaryCards(response);
    } catch (error) {
        notify.error('Failed to load summary data');
        console.error(error);
    }
}

function updateSummaryCards(data) {
    // Update total expenses
    const totalExpenses = document.getElementById('totalExpenses');
    if (totalExpenses) {
        totalExpenses.textContent = formatCurrency(data.totalExpenses);
    }

    // Update monthly spending
    const monthlySpending = document.getElementById('monthlySpending');
    if (monthlySpending) {
        monthlySpending.textContent = formatCurrency(data.monthlyTotal);
        const progress = (data.monthlyTotal / data.monthlyBudget) * 100;
        const progressBar = monthlySpending.querySelector('.progress');
        if (progressBar) {
            progressBar.style.width = `${Math.min(progress, 100)}%`;
            progressBar.style.backgroundColor = progress > 100 ? '#ef4444' : '#10b981';
        }
    }

    // Update category spending
    const categorySpending = document.getElementById('topCategories');
    if (categorySpending) {
        categorySpending.innerHTML = data.categoryTotals
            .slice(0, 3)
            .map(cat => `
                <div class="category-item">
                    <span class="category-name">${cat.name}</span>
                    <span class="category-amount">${formatCurrency(cat.total)}</span>
                </div>
            `).join('');
