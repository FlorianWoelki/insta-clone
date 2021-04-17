import React from 'react';
import Avatar from './Avatar';
import Icon from './Icon';

const Post = (): JSX.Element => {
  return (
    <div className="space-y-2">
      <img
        className="w-full cursor-pointer rounded-2xl"
        src="https://placekitten.com/300/300"
        alt="feed-image"
      />
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-2 cursor-pointer">
          <Avatar imageSize={6} />
          <p className="text-sm font-bold text-gray-800">Dom_Hill</p>
        </div>
        <div className="flex items-center space-x-3 text-gray-700">
          <div className="flex items-center space-x-1 cursor-pointer">
            <Icon name="heart" className="w-5 h-5" />
            <span className="text-sm">5.2k</span>
          </div>
          <div className="flex items-center space-x-1 cursor-pointer">
            <Icon name="chat" className="w-5 h-5" />
            <span className="text-sm">38</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Post;
