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
  onSubmit,
}) => {
  const [formData, setFormData] = useState<Omit<Category, 'id' | 'createdAt' | 'updatedAt'>>(
    initialFormData
  );
  const [newSubcategory, setNewSubcategory] = useState('');

  useEffect(() => {
    if (category) {
      setFormData({
        name: category.name,
        type: category.type,
        subcategories: [...category.subcategories],
      });
    } else {
      setFormData(initialFormData);
    }
  }, [category]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSubmit(formData);
    onClose();
  };

  const handleAddSubcategory = () => {
    if (newSubcategory.trim()) {
      setFormData({
        ...formData,
        subcategories: [...formData.subcategories, newSubcategory.trim()],
      });
      setNewSubcategory('');
    }
  };

  const handleRemoveSubcategory = (index: number) => {
    setFormData({
      ...formData,
