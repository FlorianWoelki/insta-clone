import React, { FunctionComponent } from 'react';

type AvatarProps = {
  imageSize?: Number;
  noActiveStories?: boolean;
};

const Avatar: FunctionComponent<AvatarProps> = (props): JSX.Element => {
  return (
    <div className="flex items-center justify-center">
      <div
        className={`inline-block p-0.5 rounded-full ${
          props.noActiveStories
            ? 'bg-gray-300'
            : 'bg-gradient-to-tr from-yellow-400 to-pink-600'
        }`}
      >
        <a href="#" className="relative block p-1 bg-white rounded-full">
          <img
            className={`w-${props.imageSize} h-${props.imageSize} rounded-full`}
            src="https://placekitten.com/200/200"
            alt="pb-avatar"
          ></img>

          {props.children}
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
