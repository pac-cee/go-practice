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
