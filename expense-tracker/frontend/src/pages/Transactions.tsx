import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Grid,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  TablePagination,
  Chip,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import dayjs, { Dayjs } from 'dayjs';
import { Transaction, TransactionType } from '../types';
import { useAppContext } from '../context/AppContext';
import { transactionApi } from '../services/api';
import { TransactionForm } from '../components/forms/TransactionForm';

interface FilterValues {
  startDate: Dayjs | null;
  endDate: Dayjs | null;
  type: string;
  category: string;
  account: string;
}

export const Transactions: React.FC = () => {
  const { state, dispatch } = useAppContext();
  const [openForm, setOpenForm] = useState(false);
  const [selectedTransaction, setSelectedTransaction] = useState<Transaction | null>(null);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [filters, setFilters] = useState<FilterValues>({
    startDate: dayjs().startOf('month'),
    endDate: dayjs().endOf('month'),
    type: 'all',
    category: 'all',
    account: 'all',
  });

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await transactionApi.getAll();
        dispatch({ type: 'SET_TRANSACTIONS', payload: response });
      } catch (error) {
        console.error('Error fetching transactions:', error);
      }
    };

    fetchTransactions();
  }, [dispatch]);

  const handleUpdateTransaction = async (data: Omit<Transaction, 'id' | 'createdAt' | 'updatedAt'>) => {
    if (!selectedTransaction) return;
    
    try {
      const response = await transactionApi.update(selectedTransaction.id, data);
      dispatch({ type: 'UPDATE_TRANSACTION', payload: response });
      setSelectedTransaction(null);
      setOpenForm(false);
    } catch (error) {
      console.error('Error updating transaction:', error);
    }
  };

  const handleAddTransaction = async (data: Omit<Transaction, 'id' | 'createdAt' | 'updatedAt'>) => {
    try {
      const response = await transactionApi.create(data);
      dispatch({ type: 'ADD_TRANSACTION', payload: response });
      setOpenForm(false);
    } catch (error) {
      console.error('Error adding transaction:', error);
    }
  };

  const handleDeleteTransaction = async (id: string) => {
    try {
      await transactionApi.delete(id);
      dispatch({ type: 'DELETE_TRANSACTION', payload: id });
    } catch (error) {
      console.error('Error deleting transaction:', error);
    }
  };

  const filteredTransactions = state.transactions.filter((transaction) => {
    const matchesType =
      filters.type === 'all' || transaction.type === filters.type;
    const matchesCategory =
      filters.category === 'all' || transaction.category === filters.category;
    const matchesAccount =
      filters.account === 'all' || transaction.account === filters.account;
    const matchesDate =
      (!filters.startDate || dayjs(transaction.date).isAfter(filters.startDate, 'day') || dayjs(transaction.date).isSame(filters.startDate, 'day')) &&
      (!filters.endDate || dayjs(transaction.date).isBefore(filters.endDate, 'day') || dayjs(transaction.date).isSame(filters.endDate, 'day'));

    return matchesType && matchesCategory && matchesAccount && matchesDate;
  });

