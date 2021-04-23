import React from 'react';
import Avatar from '../components/Avatar';
import { AvatarProps } from '../components/Avatar';

export default {
  title: 'UI/Avatar',
  component: Avatar,
};

export const Default = () => <Avatar />;
export const NoActiveStories = () => <Avatar noActiveStories />;
export const ImageSizeOf32 = () => <Avatar imageSize={32} />;
export const Playground = (args: AvatarProps) => <Avatar {...args} />;
