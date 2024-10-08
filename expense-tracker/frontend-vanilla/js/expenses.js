import { expenses as expensesApi } from './api.js';
import { notify, formatCurrency, formatDate, requireAuth } from './utils.js';

// State management
let expenses = [];
let currentPage = 1;
let totalPages = 1;
let filters = {
    search: '',
    category: '',
    timeRange: '',
    sortBy: 'date',
    sortOrder: 'desc'
};

// Initialize when document is ready
document.addEventListener('DOMContentLoaded', () => {
    if (!requireAuth()) return;
    initializeExpenses();
});

async function initializeExpenses() {
    // Initialize UI elements
    initializeSearchAndFilters();
    initializePagination();
    initializeExpenseForm();
    
    // Load initial data
    await loadExpenses();
}

function initializeSearchAndFilters() {
    const searchInput = document.getElementById('searchExpenses');
    const categoryFilter = document.getElementById('categoryFilter');
    const timeFilter = document.getElementById('timeFilter');

    searchInput?.addEventListener('input', debounce(async (e) => {
        filters.search = e.target.value;
