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
