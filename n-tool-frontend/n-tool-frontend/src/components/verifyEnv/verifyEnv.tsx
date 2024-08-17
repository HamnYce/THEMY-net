"use client";

import React from 'react';
import { debugLog } from '@/utils/debugLogUtil';

const VerifyEnv: React.FC = () => {
  debugLog('All Environment Variables:', process.env);

  return (
    <div>
      <h1>Verify Environment Variables</h1>
      <p>NEXT_PUBLIC_JSON_COLUMNS: {process.env.NEXT_PUBLIC_JSON_COLUMNS}</p>
      <p>NEXT_PUBLIC_JSON_FILE_PATH: {process.env.NEXT_PUBLIC_JSON_FILE_PATH}</p>
    </div>
  );
};

export default VerifyEnv;
