NAME	= monkey
compose = docker compose -p $(NAME) $(1)

run:
	$(call compose, up -d)

clean:
	$(call compose, down --remove-orphans --volumes)
