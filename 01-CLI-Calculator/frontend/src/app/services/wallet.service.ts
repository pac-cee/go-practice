import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

export interface Transaction {
  id: number;
  type: 'credit' | 'debit';
  amount: number;
  description: string;
  date: Date;
}

@Injectable({
  providedIn: 'root'
})
export class WalletService {
  private balance = new BehaviorSubject<number>(0);
  private transactions = new BehaviorSubject<Transaction[]>([]);

  constructor() {
    // Load initial data from localStorage if available
    const savedBalance = localStorage.getItem('balance');
    const savedTransactions = localStorage.getItem('transactions');
    
    if (savedBalance) {
      this.balance.next(parseFloat(savedBalance));
    }
    
    if (savedTransactions) {
      this.transactions.next(JSON.parse(savedTransactions));
    }
  }

  getBalance(): Observable<number> {
    return this.balance.asObservable();
  }

  getTransactions(): Observable<Transaction[]> {
    return this.transactions.asObservable();
