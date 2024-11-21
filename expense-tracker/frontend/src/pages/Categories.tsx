import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Grid,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Chip,
} from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { Category } from '../types';
import { useAppContext } from '../context/AppContext';
import { categoryApi } from '../services/api';
import { CategoryForm } from '../components/forms/CategoryForm';

export const Categories: React.FC = () => {
  const { state, dispatch } = useAppContext();
  const [openForm, setOpenForm] = useState(false);
  const [selectedCategory, setSelectedCategory] = useState<Category | null>(null);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const response = await categoryApi.getAll();
        dispatch({ type: 'SET_CATEGORIES', payload: response });
      } catch (error) {
        console.error('Error fetching categories:', error);
