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

