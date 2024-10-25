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
import { Account, AccountType } from '../types';

interface AccountFormProps {
  open: boolean;
  onClose: () => void;
  account?: Account | null;
  onSubmit: (data: Omit<Account, 'id' | 'createdAt' | 'updatedAt'>) => Promise<void>;
}

const initialFormData: Omit<Account, 'id' | 'createdAt' | 'updatedAt'> = {
  name: '',
  type: 'checking',
  balance: 0,
  currency: 'USD',
};

export const AccountForm: React.FC<AccountFormProps> = ({
  open,
  onClose,
  account,
  onSubmit,
}) => {
  const [formData, setFormData] = useState<Omit<Account, 'id' | 'createdAt' | 'updatedAt'>>(initialFormData);

  useEffect(() => {
    if (account) {
      setFormData({
        name: account.name,
        type: account.type,
        balance: account.balance,
        currency: account.currency,
      });
    } else {
      setFormData(initialFormData);
    }
  }, [account]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSubmit(formData);
    onClose();
  };
