import React, {useEffect, useState} from 'react';
import {pingApi} from './api';
import {ThemeProvider} from "./components/theme-provider.tsx";
import {ToastProvider} from "@/components/ui/toast.tsx";
import {Toaster} from "@/components/ui/toaster.tsx";
import Logo from './assets/logo.png';
import {Separator} from "@/components/ui/separator.tsx";

const App: React.FC = () => {
  const [apiStatus, setApiStatus] = useState<boolean>(false);
  useEffect(() => {
    pingApi().then((_) => setApiStatus(true));
  }, []);

  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <ToastProvider>
        <div className="text-white min-h-screen">
          <div className="container mx-auto p-5">
            <div className="mb-5 flex">
              <img src={Logo} alt="logo" className="h-12"/>
              <span className="flex-auto-leading-none ml-2">
                {/* add version */}
              </span>
            </div>
            <Separator/>
            <p className="text-xl font-bold mt-10 mb-10">
              API Status: {apiStatus ? 'OK' : 'NOT OK'}
            </p>
          </div>
        </div>
        <Toaster />
      </ToastProvider>
    </ThemeProvider>
  );
};

export default App;
