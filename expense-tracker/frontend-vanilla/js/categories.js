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
        const colors = ['#ef4444', '#f59e0b', '#10b981', '#3b82f6', '#8b5cf6', '#ec4899'];
        colorPicker.innerHTML = colors.map(color => `
            <div class="color-option" style="background-color: ${color}" 
                 data-color="${color}" onclick="selectColor('${color}')"></div>
        `).join('');
    }

    // Initialize icon picker
    if (iconPicker) {
        const icons = ['utensils', 'car', 'bolt', 'film', 'shopping-bag', 'heart', 
                      'home', 'plane', 'gift', 'book', 'graduation-cap', 'briefcase'];
        iconPicker.innerHTML = icons.map(icon => `
            <div class="icon-option" data-icon="fa-${icon}" onclick="selectIcon('fa-${icon}')">
                <i class="fas fa-${icon}"></i>
            </div>
        `).join('');
    }

    addCategoryBtn?.addEventListener('click', () => {
        currentCategory = null;
        categoryForm.reset();
        categoryModal.classList.remove('hidden');
    });

    cancelCategoryBtn?.addEventListener('click', () => {
        categoryModal.classList.add('hidden');
    });

    categoryForm?.addEventListener('submit', handleCategorySubmit);
}

async function loadCategories() {
    try {
        const response = await categoriesApi.getAll();
        categories = response.categories;
        
        if (categories.length === 0) {
            // For new users, create default categories
