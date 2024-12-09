export type AccountType = 'bank' | 'cash' | 'mobile_money';
export type TransactionType = 'income' | 'expense';
export type BudgetPeriod = 'monthly' | 'yearly' | 'custom';

export interface BaseModel {
  id: string;
  createdAt: string;
  updatedAt: string;
}

export interface Account extends BaseModel {
  name: string;
  type: AccountType;
  balance: number;
  currency: string;
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
  type: TransactionType;
  subcategories: string[];
}

export interface Budget extends BaseModel {
  amount: number;
  spent: number;
  category: string;
  period: BudgetPeriod;
  startDate: string;
  endDate: string;
}

export interface TransactionReport {
  totalIncome: number;
  totalExpenses: number;
  netAmount: number;
  transactions: Transaction[];
  categories: CategoryExpenseSummary[];
}

export interface CategoryExpenseSummary {
  category: string;
  amount: number;
  count: number;
}

export interface BudgetReport {
  category: string;
  budgetAmount: number;
  spentAmount: number;
  percentage: number;
  startDate: string;
  endDate: string;
}

export interface AppState {
  transactions: Transaction[];
  categories: Category[];
  budgets: Budget[];
  accounts: Account[];
  loading: boolean;
  error: string | null;
}

export type AppAction =
  | { type: 'SET_TRANSACTIONS'; payload: Transaction[] }
  | { type: 'ADD_TRANSACTION'; payload: Transaction }
  | { type: 'UPDATE_TRANSACTION'; payload: Transaction }
  | { type: 'DELETE_TRANSACTION'; payload: string }
  | { type: 'SET_CATEGORIES'; payload: Category[] }
  | { type: 'ADD_CATEGORY'; payload: Category }
  | { type: 'UPDATE_CATEGORY'; payload: Category }
  | { type: 'DELETE_CATEGORY'; payload: string }
  | { type: 'SET_BUDGETS'; payload: Budget[] }
  | { type: 'ADD_BUDGET'; payload: Budget }
  | { type: 'UPDATE_BUDGET'; payload: Budget }
  | { type: 'DELETE_BUDGET'; payload: string }
  | { type: 'SET_ACCOUNTS'; payload: Account[] }
  | { type: 'ADD_ACCOUNT'; payload: Account }
  | { type: 'UPDATE_ACCOUNT'; payload: Account }
  | { type: 'DELETE_ACCOUNT'; payload: string }
  | { type: 'SET_LOADING'; payload: boolean }
  | { type: 'SET_ERROR'; payload: string | null };

export interface ChartData {
  labels: string[];
  datasets: {
    label: string;
    data: number[];
    backgroundColor?: string | string[];
    borderColor?: string | string[];
    borderWidth?: number;
    fill?: boolean;
  }[];
}

export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export interface ApiError {
  message: string;
  errors?: Record<string, string>;
}
