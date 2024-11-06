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
import { Budget, Category } from '../../types';
import { useAppContext } from '../../context/AppContext';

interface BudgetFormProps {
  open: boolean;
  onClose: () => void;
  budget?: Budget | null;
  onSubmit: (data: Omit<Budget, 'id' | 'createdAt' | 'updatedAt' | 'spent'>) => Promise<void>;
}

const initialFormData: Omit<Budget, 'id' | 'createdAt' | 'updatedAt' | 'spent'> = {
  amount: 0,
  category: '',
  startDate: '',
  endDate: '',
};

export const BudgetForm: React.FC<BudgetFormProps> = ({
  open,
  onClose,
  budget,
  onSubmit,
}) => {
  const { state } = useAppContext();
  const [formData, setFormData] = useState<Omit<Budget, 'id' | 'createdAt' | 'updatedAt' | 'spent'>>(
    initialFormData
  );

  useEffect(() => {
    if (budget) {
      setFormData({
        amount: budget.amount,
        category: budget.category,
        startDate: budget.startDate,
        endDate: budget.endDate,
      });
    } else {
      setFormData(initialFormData);
    }
  }, [budget]);
