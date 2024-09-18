// API endpoints and fetch utilities
const API_URL = 'http://localhost:8080/api';

// Authentication API calls
const auth = {
    async login(email, password) {
        try {
            const response = await fetch(`${API_URL}/auth/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password }),
            });
            const data = await response.json();
            if (!response.ok) throw new Error(data.message);
            return data;
        } catch (error) {
            throw error;
        }
    },

    async register(userData) {
        try {
            const response = await fetch(`${API_URL}/auth/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData),
            });
            const data = await response.json();
            if (!response.ok) throw new Error(data.message);
            return data;
        } catch (error) {
            throw error;
        }
    },

    async logout() {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/pages/login.html';
    }
};

// Expenses API calls
const expenses = {
    async getAll(params = {}) {
        try {
            const queryString = new URLSearchParams(params).toString();
            const response = await fetch(`${API_URL}/expenses?${queryString}`, {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`,
                }
            });
            const data = await response.json();
            if (!response.ok) throw new Error(data.message);
            return data;
        } catch (error) {
            throw error;
        }
    },

    async create(expenseData) {
        try {
            const response = await fetch(`${API_URL}/expenses`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`,
                },
                body: JSON.stringify(expenseData),
            });
            const data = await response.json();
            if (!response.ok) throw new Error(data.message);
