set PATH=%PATH%;C:\Program Files\nodejs;C:\Users\Shira\AppData\Roaming\npm
call npm.cmd install -g yarn
call yarn.cmd create vite frontend --template react
cd frontend
call yarn.cmd install
call yarn.cmd add @xyflow/react lucide-react tailwindcss @tailwindcss/vite
