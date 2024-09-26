import { categories as categoriesApi } from './api.js';
import { notify, requireAuth } from './utils.js';

// State management
let categories = [];
let currentCategory = null;

// Default categories (used only for new users)
const defaultCategories = [
    { name: 'Food & Dining', icon: 'fa-utensils', color: '#ef4444', budget: 500 },
    { name: 'Transportation', icon: 'fa-car', color: '#3b82f6', budget: 300 },
    { name: 'Utilities', icon: 'fa-bolt', color: '#f59e0b', budget: 200 },
    { name: 'Entertainment', icon: 'fa-film', color: '#8b5cf6', budget: 150 },
    { name: 'Shopping', icon: 'fa-shopping-bag', color: '#ec4899', budget: 400 },
    { name: 'Healthcare', icon: 'fa-heart', color: '#10b981', budget: 200 }
];

// Initialize when document is ready
document.addEventListener('DOMContentLoaded', () => {
    if (!requireAuth()) return;
    initializeCategories();
});

async function initializeCategories() {
    initializeCategoryForm();
    await loadCategories();
}

function initializeCategoryForm() {
    const addCategoryBtn = document.getElementById('addCategoryBtn');
    const categoryModal = document.getElementById('categoryModal');
    const categoryForm = document.getElementById('categoryForm');
    const cancelCategoryBtn = document.getElementById('cancelCategory');
    const colorPicker = document.getElementById('categoryColor');
    const iconPicker = document.getElementById('categoryIcon');

    // Initialize color picker
    if (colorPicker) {
