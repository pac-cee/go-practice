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
