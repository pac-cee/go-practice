import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Grid,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Chip,
} from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { Account, AccountType } from '../types';
import { useAppContext } from '../context/AppContext';
import { accountApi } from '../services/api';
import { AccountForm } from '../components/forms/AccountForm';

export const Accounts: React.FC = () => {
  const { state, dispatch } = useAppContext();
  const [openForm, setOpenForm] = useState(false);
  const [selectedAccount, setSelectedAccount] = useState<Account | null>(null);

  useEffect(() => {
    const fetchAccounts = async () => {
      try {
        const response = await accountApi.getAll();
        dispatch({ type: 'SET_ACCOUNTS', payload: response });
      } catch (error) {
        console.error('Error fetching accounts:', error);
      }
    };

    fetchAccounts();
  }, [dispatch]);

  const handleAddAccount = async (data: Omit<Account, 'id' | 'createdAt' | 'updatedAt'>) => {
    try {
      const response = await accountApi.create(data);
      dispatch({ type: 'ADD_ACCOUNT', payload: response });
      setOpenForm(false);
    } catch (error) {
      console.error('Error adding account:', error);
    }
  };

  const handleUpdateAccount = async (data: Omit<Account, 'id' | 'createdAt' | 'updatedAt'>) => {
    if (!selectedAccount) return;
    
    try {
      const response = await accountApi.update(selectedAccount.id, data);
      dispatch({ type: 'UPDATE_ACCOUNT', payload: response });
      setSelectedAccount(null);
      setOpenForm(false);
    } catch (error) {
      console.error('Error updating account:', error);
    }
  };

  const handleDeleteAccount = async (id: string) => {
    try {
      await accountApi.delete(id);
      dispatch({ type: 'DELETE_ACCOUNT', payload: id });
    } catch (error) {
      console.error('Error deleting account:', error);
    }
  };

