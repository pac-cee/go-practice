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

