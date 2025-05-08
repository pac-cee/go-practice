import React, { useState } from 'react';

interface ActivityFormProps {
  onSubmit: (activity: { title: string; description: string; duration: number }) => void;
}

export default function ActivityForm({ onSubmit }: ActivityFormProps) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [duration, setDuration] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      title,
      description,
      duration: parseInt(duration),
    });
    setTitle('');
