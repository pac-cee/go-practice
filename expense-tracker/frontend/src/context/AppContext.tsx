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

export const useAppContext = () => useContext(AppContext);

const appReducer = (state: AppState, action: AppAction): AppState => {
  switch (action.type) {
    case 'SET_TRANSACTIONS':
      return {
        ...state,
        transactions: action.payload,
      };

    case 'ADD_TRANSACTION':
      return {
        ...state,
        transactions: [...state.transactions, action.payload],
      };

    case 'UPDATE_TRANSACTION':
      return {
