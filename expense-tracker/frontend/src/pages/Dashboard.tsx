import React, { useState, useEffect } from 'react';
import {
  Grid,
  Paper,
  Typography,
  Box,
  useTheme,
} from '@mui/material';
import { styled } from '@mui/material/styles';

import { useAppContext } from '../context/AppContext';
import { reportApi } from '../services/api';
import ExpensePieChart from '../components/charts/ExpensePieChart';
import { IncomeExpenseChart } from '../components/charts/IncomeExpenseChart';
import { TrendLineChart } from '../components/charts/TrendLineChart';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  height: '100%',
}));

const Dashboard: React.FC = () => {
  const theme = useTheme();
  const { state } = useAppContext();
  const [totalIncome, setTotalIncome] = useState<number>(0);
  const [totalExpenses, setTotalExpenses] = useState<number>(0);
  const [netAmount, setNetAmount] = useState<number>(0);
  const [categoryExpenses, setCategoryExpenses] = useState<any[]>([]);
  const [monthlyData, setMonthlyData] = useState<any[]>([]);

  useEffect(() => {
    const fetchDashboardData = async () => {
      try {
        const startDate = new Date();
        startDate.setMonth(startDate.getMonth() - 1);
        
        const [transactionReport, categoryReport] = await Promise.all([
          reportApi.getTransactionReport({ startDate: startDate.toISOString() }),
