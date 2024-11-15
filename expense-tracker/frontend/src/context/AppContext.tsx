import React, { createContext, useContext, useReducer } from 'react';
import { Transaction, Category, Budget, Account, AppState, AppAction } from '../types';

const initialState: AppState = {
  transactions: [],
  categories: [],
  budgets: [],
  accounts: [],
  loading: false,
  error: null,
};

const AppContext = createContext<{
  state: AppState;
  dispatch: React.Dispatch<AppAction>;
}>({
  state: initialState,
  dispatch: () => null,
});
