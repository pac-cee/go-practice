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
