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
