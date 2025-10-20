@echo off
REM ==============================================
REM Staking 时间控制测试脚本
REM 适用于 Windows 测试人员
REM ==============================================

echo.
echo ============================================
echo Staking 时间控制测试工具
echo ============================================
echo.
echo 请选择操作:
echo.
echo 1. 查看当前时间状态
echo 2. 启用测试模式
echo 3. 快进时间 (分钟)
echo 4. 快进时间 (天)
echo 5. 禁用测试模式
echo 6. 验证环境配置
echo 7. 退出
echo.
set /p choice=请输入选项 (1-7): 

if "%choice%"=="1" goto check_status
if "%choice%"=="2" goto enable_test
if "%choice%"=="3" goto forward_minutes
if "%choice%"=="4" goto forward_days
if "%choice%"=="5" goto disable_test
if "%choice%"=="6" goto verify
if "%choice%"=="7" goto end

echo 无效选项，请重新选择
goto menu

:check_status
echo.
echo 正在查看时间状态...
npx hardhat run scripts\check-time-status.ts --network sepoliaCustom
goto menu

:enable_test
echo.
echo 正在启用测试模式...
npx hardhat run scripts\enable-test-mode.ts --network sepoliaCustom
echo.
echo 测试模式已启用！
goto menu

:forward_minutes
echo.
set /p minutes=请输入要快进的分钟数: 
echo 正在快进 %minutes% 分钟...
npx hardhat run scripts\fast-forward-time.ts --network sepoliaCustom -- --minutes %minutes%
echo.
echo 时间已快进！
goto menu

:forward_days
echo.
set /p days=请输入要快进的天数: 
echo 正在快进 %days% 天...
npx hardhat run scripts\fast-forward-time.ts --network sepoliaCustom -- --days %days%
echo.
echo 时间已快进！
goto menu

:disable_test
echo.
echo 正在禁用测试模式...
npx hardhat run scripts\disable-test-mode.ts --network sepoliaCustom
echo.
echo 测试模式已禁用！
goto menu

:verify
echo.
echo 正在验证环境配置...
npx hardhat run scripts\check-staking-owner.ts --network sepoliaCustom
goto menu

:menu
echo.
echo ============================================
echo.
set /p continue=是否继续? (Y/N): 
if /i "%continue%"=="Y" goto start
if /i "%continue%"=="N" goto end
goto menu

:start
cls
goto top

:top
goto :eof

:end
echo.
echo 感谢使用！
echo.
pause

