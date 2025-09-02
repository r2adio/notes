-- here Car can be used as a class
local function Car(color)
	-- here, private
	return {
		-- here, public
		color = color or "black",

		runs_on = function() -- `.` syntax is used.
			print("road")
		end,

		fuel = function(self) -- `:` syntax is used.
			-- self is used to refer the table itself.
			-- thus self can be used to reference anything inside this table.

			print("petrol") -- this works fine, when using `.`
			print("petrol", self.color) -- this gives err with `.`, thus use `:` syntax.
			-- above statement give err, 'cause self is nill value.
			-- 'cause using `:` syntax means, u're passing obj into itself. _jeep:fuel(jeep)_
		end,
	}
end

-- making a object of class Car
local jeep = Car()
local suv = Car("red")

print(jeep.color)
print(suv.color)
jeep.runs_on()
jeep.fuel(jeep) -- statement gives err, if the obj itself is not passed into itself.
jeep:fuel() -- another alternative, syntax sugar, uses `:` instead of `.` when using self in function signature.
