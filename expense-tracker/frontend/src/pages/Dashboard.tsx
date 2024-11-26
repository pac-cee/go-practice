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
