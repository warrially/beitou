math.randomseed(os.time())


-- local total = 0


-- for j = 1, 20 do
--     local n = 10000
--     local bet = 1

--     for i=1,1000 do
--         if n < bet then
--             bet = 1            
--         end;

--         if n <= 0 then
--             -- print("玩家没钱")
--             break
--         end
 

--         local a = math.random(13)
--         local b = math.random(13)

--         if a > b then
--             -- 赢钱
--             n = n + bet * 0.97
--             bet = 1
--         elseif a < b then
--             -- 输钱
--             n = n - bet
--             bet = bet * 2 
--         end

--         if bet > 1024 then
--          --   print("玩家的钱", n, "下注额:=",  bet, "次数:= ", i)
--         end 
--     end;
--     total = total + n - 10000
--     print("第" .. j.. "个玩家的钱是" ..  n, (n > 10000 and "赢" or "输"))
-- end;


-- print("总赢", total)


for j = 1, 100 do
    local n = 0
    for i=1,48 do
        if math.random(100) <= 3 then
            n = n + 1
        end
    end
    print(n)
end
