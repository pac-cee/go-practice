import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Grid,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  TablePagination,
  Chip,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import dayjs, { Dayjs } from 'dayjs';
import { Transaction, TransactionType } from '../types';
import { useAppContext } from '../context/AppContext';
import { transactionApi } from '../services/api';
import { TransactionForm } from '../components/forms/TransactionForm';

interface FilterValues {
  startDate: Dayjs | null;
  endDate: Dayjs | null;
  type: string;
  category: string;
