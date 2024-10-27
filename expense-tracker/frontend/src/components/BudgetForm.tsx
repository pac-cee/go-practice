import { useState, useEffect } from 'react';
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
  Grid,
  InputAdornment,
} from '@mui/material';
import { Budget } from '../types';
import { useAppContext } from '../context/AppContext';
import { budgetApi } from '../services/api';

interface BudgetFormProps {
  open: boolean;
  onClose: () => void;
  budget?: Budget;
}

const BudgetForm = ({ open, onClose, budget }: BudgetFormProps) => {
  const { state, dispatch } = useAppContext();
  const [formData, setFormData] = useState({
    category: '',
    amount: '',
    period: 'monthly',
  });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (budget) {
      setFormData({
        category: budget.category,
        amount: String(budget.amount),
        period: budget.period,
      });
    }
  }, [budget]);

  const handleChange = (field: string) => (event: any) => {
    setFormData((prev) => ({
      ...prev,
      [field]: event.target.value,
    }));
  };

  const handleSubmit = async () => {
    if (!formData.category || !formData.amount) return;

    try {
      setLoading(true);
      const data = {
