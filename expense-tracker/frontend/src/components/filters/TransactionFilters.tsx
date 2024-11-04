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
      <Box sx={{ display: 'flex', gap: 2 }}>
        <DatePicker
          label="Start Date"
          value={filters.startDate ? dayjs(filters.startDate) : null}
          onChange={(date: Dayjs | null) =>
            handleFilterChange('startDate', date?.toISOString())
          }
          slotProps={{
            textField: {
              fullWidth: true,
            },
          }}
        />

        <DatePicker
          label="End Date"
          value={filters.endDate ? dayjs(filters.endDate) : null}
          onChange={(date: Dayjs | null) =>
            handleFilterChange('endDate', date?.toISOString())
          }
          slotProps={{
            textField: {
              fullWidth: true,
            },
          }}
        />
      </Box>

      <Box sx={{ display: 'flex', gap: 2 }}>
        <FormControl fullWidth>
          <InputLabel>Type</InputLabel>
          <Select
            value={filters.type || ''}
            onChange={(e) =>
              handleFilterChange('type', e.target.value as TransactionType)
            }
            label="Type"
          >
            <MenuItem value="">All</MenuItem>
            <MenuItem value="expense">Expense</MenuItem>
            <MenuItem value="income">Income</MenuItem>
          </Select>
        </FormControl>

        <FormControl fullWidth>
          <InputLabel>Category</InputLabel>
          <Select
            value={filters.category || ''}
            onChange={(e) => handleFilterChange('category', e.target.value)}
            label="Category"
          >
            <MenuItem value="">All</MenuItem>
            {state.categories
              .filter((category: Category) =>
                filters.type ? category.type === filters.type : true
              )
              .map((category: Category) => (
