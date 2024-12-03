import axios from 'axios';
import { Transaction, Category, Budget, Account, TransactionReport, BudgetReport, CategoryExpenseSummary } from '../types';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Transaction API
export const transactionApi = {
  getAll: async () => {
    const response = await api.get<Transaction[]>('/transactions');
    return response.data;
  },

  getById: async (id: string) => {
    const response = await api.get<Transaction>(`/transactions/${id}`);
    return response.data;
  },

  create: async (transaction: Omit<Transaction, 'id' | 'createdAt' | 'updatedAt'>) => {
    const response = await api.post<Transaction>('/transactions', transaction);
    return response.data;
  },

  update: async (id: string, transaction: Partial<Transaction>) => {
    const response = await api.put<Transaction>(`/transactions/${id}`, transaction);
    return response.data;
  },

  delete: async (id: string) => {
    await api.delete(`/transactions/${id}`);
  },
};

// Category API
export const categoryApi = {
  getAll: async () => {
    const response = await api.get<Category[]>('/categories');
    return response.data;
  },

  getById: async (id: string) => {
    const response = await api.get<Category>(`/categories/${id}`);
    return response.data;
  },

  create: async (category: Omit<Category, 'id' | 'createdAt' | 'updatedAt'>) => {
    const response = await api.post<Category>('/categories', category);
    return response.data;
  },

  update: async (id: string, category: Partial<Category>) => {
    const response = await api.put<Category>(`/categories/${id}`, category);
    return response.data;
  },

  delete: async (id: string) => {
    await api.delete(`/categories/${id}`);
  },
};

// Budget API
export const budgetApi = {
  getAll: async () => {
    const response = await api.get<Budget[]>('/budgets');
    return response.data;
  },

  getById: async (id: string) => {
    const response = await api.get<Budget>(`/budgets/${id}`);
    return response.data;
  },

  create: async (budget: Omit<Budget, 'id' | 'spent' | 'createdAt' | 'updatedAt'>) => {
    const response = await api.post<Budget>('/budgets', budget);
    return response.data;
  },

  update: async (id: string, budget: Partial<Budget>) => {
    const response = await api.put<Budget>(`/budgets/${id}`, budget);
    return response.data;
  },

  delete: async (id: string) => {
    await api.delete(`/budgets/${id}`);
  },
};

// Account API
export const accountApi = {
