-- inheritance: using the props of one class in another

local function Pet(name)
	return {
		name = name or "charlie",
		status = "hungry",

		speak = function(self)
			print("meow")
		end,
	}
end

local function Dog(name, breed)
	-- inheriting the Pet class into Dog
	local dog = Pet(name)

	-- adding a new property to the dog class
	dog.breed = breed or "doberman"
	dog.loyalty = 0

	-- over-writing the function
	dog.speak = function(self)
		print(self.name .. " says: " .. "woof")
	end

	return dog
end

local doberman = Dog("Jesse", "poodle")
print(doberman.name, doberman.breed)
doberman:speak() -- used the `:` to call the function, even though it's not using the self in function. a better practice
