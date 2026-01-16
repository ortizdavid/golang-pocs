c

func (h *ProductHandler) update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product models.ProductModel
	c.BodyParser(&product)
	product.ID = id

	if err := h.repository.Update(product); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": i18n.T(c, "product.not_found"),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": i18n.T(c, "product.updated"),
	})
}

func (h *ProductHandler) delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.repository.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": i18n.T(c, "product.error.delete"),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": i18n.T(c, "product.delete"),
	})
}

func (h *ProductHandler) getByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product, err := h.repository.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": i18n.T(c, "product.not_found"),
		})
	}

	return c.Status(200).JSON(product)
}