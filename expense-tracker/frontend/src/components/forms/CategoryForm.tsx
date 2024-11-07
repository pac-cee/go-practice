import React, { useState, useEffect } from 'react';
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
  Box,
  IconButton,
  List,
  ListItem,
  ListItemText,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import { Category, CategoryType } from '../../types';

interface CategoryFormProps {
  open: boolean;
  onClose: () => void;
  category?: Category | null;
  onSubmit: (data: Omit<Category, 'id' | 'createdAt' | 'updatedAt'>) => Promise<void>;
}

const initialFormData: Omit<Category, 'id' | 'createdAt' | 'updatedAt'> = {
  name: '',
  type: 'expense',
  subcategories: [],
};

export const CategoryForm: React.FC<CategoryFormProps> = ({
  open,
  onClose,
  category,
