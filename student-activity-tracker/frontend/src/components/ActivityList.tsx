import React from 'react';

interface Activity {
  id: string;
  title: string;
  description: string;
  duration: number;
  status: string;
  startTime?: string;
  endTime?: string;
}

interface ActivityListProps {
  activities: Activity[];
  onStart: (id: string) => void;
  onComplete: (id: string) => void;
  onDelete: (id: string) => void;
}

export default function ActivityList({ activities, onStart, onComplete, onDelete }: ActivityListProps) {
  const getStatusColor = (status: string) => {
    switch (status) {
      case 'planned':
        return 'bg-yellow-100 text-yellow-800';
      case 'in-progress':
        return 'bg-blue-100 text-blue-800';
      case 'completed':
        return 'bg-green-100 text-green-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div className="space-y-4">
      {activities.map((activity) => (
        <div
          key={activity.id}
