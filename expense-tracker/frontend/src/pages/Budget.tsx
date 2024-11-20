import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Grid,
  LinearProgress,
} from '@mui/material';
import { BudgetForm } from '../components/forms/BudgetForm';
import { Budget } from '../types';
import { useAppContext } from '../context/AppContext';

export const BudgetPage: React.FC = () => {
  const { state, dispatch } = useAppContext();
  const [openForm, setOpenForm] = useState(false);
  const [selectedBudget, setSelectedBudget] = useState<Budget | null>(null);

  const handleOpenForm = (budget?: Budget) => {
    setSelectedBudget(budget || null);
    setOpenForm(true);
  };

  const handleCloseForm = () => {
    setSelectedBudget(null);
    setOpenForm(false);
  };

  const handleSubmit = async (data: Omit<Budget, 'id' | 'createdAt' | 'updatedAt' | 'spent'>) => {
    try {
      if (selectedBudget) {
        // Update existing budget
        const response = await fetch(`/api/budgets/${selectedBudget.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });

        if (!response.ok) {
          throw new Error('Failed to update budget');
        }

        const updatedBudget = await response.json();
        dispatch({ type: 'UPDATE_BUDGET', payload: updatedBudget });
      } else {
        // Create new budget
        const response = await fetch('/api/budgets', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });

        if (!response.ok) {
          throw new Error('Failed to create budget');
        }

        const newBudget = await response.json();
        dispatch({ type: 'ADD_BUDGET', payload: newBudget });
      }

      handleCloseForm();
    } catch (error) {
      console.error('Error submitting budget:', error);
    }
  };

  const handleDelete = async (id: string) => {
    try {
      const response = await fetch(`/api/budgets/${id}`, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error('Failed to delete budget');
      }

      dispatch({ type: 'DELETE_BUDGET', payload: id });
    } catch (error) {
      console.error('Error deleting budget:', error);
    }
  };

  useEffect(() => {
    const fetchBudgets = async () => {
      try {
        const response = await fetch('/api/budgets');
        if (!response.ok) {
          throw new Error('Failed to fetch budgets');
        }

        const budgets = await response.json();
        dispatch({ type: 'SET_BUDGETS', payload: budgets });
      } catch (error) {
        console.error('Error fetching budgets:', error);
      }
    };

    fetchBudgets();
  }, [dispatch]);

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
        <Typography variant="h4">Budgets</Typography>
        <Button variant="contained" color="primary" onClick={() => handleOpenForm()}>
          Add Budget
        </Button>
      </Box>

