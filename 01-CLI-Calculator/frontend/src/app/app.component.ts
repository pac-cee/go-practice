import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  balance: number = 0;
  transactionType: string = 'credit';
  amount: number = 0;
  description: string = '';
  transactions: Transaction[] = [];

  addTransaction() {
    if (this.amount <= 0 || !this.description) {
      return;
    }

