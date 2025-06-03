package application

//func TestOrderService_Create(t *testing.T) {
//	conn, _ := db.PostgresConn()
//	partnerRepo := db.NewPartnerRepository(conn)
//	eventRepo := db.NewEventRepository(conn)
//	orderRepo := db.NewOrderRepository(conn)
//	customerRepo := db.NewCustomerRepository(conn)
//	customer, _ := entity.CreateCustomer("443.376.887-14", "Danilo Bandeira")
//	_ = customerRepo.Save(customer)
//	uow := unitofwork.NewUoW(conn)
//	uow.RegisterFactory("PartnerRepository", func(exec db.Executor) any {
//		return db.NewPartnerRepository(exec)
//	})
//	uow.RegisterFactory("EventRepository", func(exec db.Executor) any {
//		return db.NewEventRepository(exec)
//	})
//	uow.RegisterFactory("OrderRepository", func(exec db.Executor) any {
//		return db.NewOrderRepository(exec)
//	})
//	uow.RegisterFactory("SpotReservationRepository", func(exec db.Executor) any {
//		return db.NewSpotReservationRepository(exec)
//	})
//	partner, _ := entity.CreatePartner("Partner 1", time.Now())
//	_ = partnerRepo.Save(partner)
//	event, _ := partner.CreateEvent(entity.PartnerCreateEvent{
//		Name:        "OrderCreate" + time.Now().Format(time.RFC3339),
//		Description: nil,
//		Date:        time.Now(),
//	})
//	_ = event.AddSection(entity.AddSectionCommand{
//		Name:        "Premium",
//		Description: nil,
//		TotalSpots:  100,
//		Price:       1999.99,
//	})
//	var sectionID entity.EventSectionID
//	var spotID entity.EventSpotID
//	for _, s := range event.Sections.Data {
//		sectionID = s.ID
//		for _, spot := range s.Spots.Data {
//			spotID = spot.ID
//			break
//		}
//		break
//	}
//	_ = eventRepo.Save(event)
//	service := NewOrderService(orderRepo, uow)
//	order, err := service.Create(OrderCreateInput{
//		EventID:    event.ID.String(),
//		SectionID:  sectionID.String(),
//		SpotID:     spotID.String(),
//		CustomerID: customer.ID.String(),
//	})
//	if err != nil {
//		t.Errorf("expected error to be empty\ngot: %v", err)
//		return
//	}
//	o, err := orderRepo.FindByID(order.ID.String())
//	if err != nil {
//		t.Errorf("expected error to be empty\ngot: %v", err)
//		return
//	}
//	if !order.ID.Equal(o.ID) {
//		t.Errorf("expected order not created\nexpected id: %s\ngot: %s", order.ID.String(), o.ID.String())
//		return
//	}
//}
