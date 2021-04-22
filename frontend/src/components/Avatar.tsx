import React, { FunctionComponent } from 'react';

export type AvatarProps = {
  imageSize?: Number;
  noActiveStories?: boolean;
};

const Avatar: FunctionComponent<AvatarProps> = ({
  noActiveStories,
  imageSize,
}): JSX.Element => {
  return (
    <div className="flex items-center justify-center">
      <div
        className={`inline-block p-0.5 rounded-full ${
          noActiveStories
            ? 'bg-gray-300'
            : 'bg-gradient-to-tr from-yellow-400 to-pink-600'
        }`}>
        <a href="#" className="relative block p-1 bg-white rounded-full">
          <img
            className={`w-${imageSize} h-${imageSize} rounded-full`}
            src="https://placekitten.com/200/200"
            alt="pb-avatar"></img>
        </a>
      </div>
    </div>
  );
};

Avatar.defaultProps = {
  imageSize: 16,
  noActiveStories: false,
};

export default Avatar;
