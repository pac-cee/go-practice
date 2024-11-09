import React, { useState, useEffect } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Box,
  InputAdornment,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import dayjs, { Dayjs } from 'dayjs';
import { Transaction, TransactionType, Category, Account } from '../../types';
import { useAppContext } from '../../context/AppContext';

interface TransactionFormProps {
  open: boolean;
  onClose: () => void;
  transaction?: Transaction | null;
  onSubmit: (data: Omit<Transaction, 'id' | 'createdAt' | 'updatedAt'>) => Promise<void>;
}

const initialFormData: Omit<Transaction, 'id' | 'createdAt' | 'updatedAt'> = {
  amount: 0,
  type: 'expense',
  category: '',
  subcategory: '',
  account: '',
  description: '',
  date: dayjs().toISOString(),
};

export const TransactionForm: React.FC<TransactionFormProps> = ({
