import React from 'react';
import {
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  InputAdornment,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import dayjs, { Dayjs } from 'dayjs';
import { TransactionType, TransactionFilters, Category, Account } from '../../types';
import { useAppContext } from '../../context/AppContext';

interface TransactionFiltersProps {
  filters: TransactionFilters;
  onFilterChange: (filters: TransactionFilters) => void;
}

export const TransactionFiltersComponent: React.FC<TransactionFiltersProps> = ({
  filters,
  onFilterChange,
}) => {
  const { state } = useAppContext();

  const handleFilterChange = (
    field: keyof TransactionFilters,
    value: string | number | undefined
  ) => {
    onFilterChange({
      ...filters,
      [field]: value,
    });
  };

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mb: 3 }}>
