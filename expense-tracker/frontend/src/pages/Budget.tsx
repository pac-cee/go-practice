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
