export type TransactionType = 'income' | 'expense';
export type AccountType = 'checking' | 'savings' | 'credit' | 'cash' | 'investment';
export type BudgetPeriod = 'monthly' | 'yearly' | 'custom';
export type CategoryType = 'income' | 'expense';

export interface BaseModel {
  id: string;
  createdAt: string;
  updatedAt: string;
}

export interface Transaction extends BaseModel {
  amount: number;
  type: TransactionType;
  category: string;
  subcategory?: string;
  account: string;
  description: string;
  date: string;
}

export interface Category extends BaseModel {
  name: string;
  type: CategoryType;
  subcategories: string[];
}

export interface Budget extends BaseModel {
  amount: number;
  spent: number;
  category: string;
  startDate: string;
  endDate: string;
}

export interface Account extends BaseModel {
  name: string;
  type: AccountType;
  balance: number;
  currency: string;
}

export interface CategoryExpenseSummary {
  category: string;
  amount: number;
  percentage: number;
}

export interface TransactionFilters {
  startDate?: string;
  endDate?: string;
  type?: TransactionType;
  category?: string;
  account?: string;
  minAmount?: number;
  maxAmount?: number;
}

export interface TransactionReport {
  totalIncome: number;
  totalExpenses: number;
  netIncome: number;
  categoryBreakdown: CategoryExpenseSummary[];
}

export interface BudgetReport {
  totalBudget: number;
  totalSpent: number;
  remaining: number;
  categories: {
    category: string;
    budget: number;
    spent: number;
    remaining: number;
    percentage: number;
  }[];
